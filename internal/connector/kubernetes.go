package connector

import (
	"context"
	"strings"

	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Kubernetes struct {
	Config *rest.Config
}

func NewKubernetesConnector(url string, ca string, token string) Kubernetes {
	return Kubernetes{
		Config: &rest.Config{
			Host:        url,
			BearerToken: token,
			TLSClientConfig: rest.TLSClientConfig{
				CAData: []byte(ca),
			},
		},
	}
}

func normalize(s string) string {
	return strings.ReplaceAll(s, "@", "-")
}

func (k *Kubernetes) CreateToken(user string) (string, error) {
	// todo: create a new token vs returning an existing one
	// todo: invalidate default token and generate new ones?
	clientset, err := kubernetes.NewForConfig(k.Config)
	if err != nil {
		return "", err
	}

	name := normalize(user)

	sa, err := clientset.CoreV1().ServiceAccounts("default").Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	var data string
	for _, s := range sa.Secrets {
		secret, err := clientset.CoreV1().Secrets("default").Get(context.Background(), s.Name, metav1.GetOptions{})
		if err != nil {
			return "", err
		}

		data = string(secret.Data["token"])
		break
	}

	return data, nil
}

func (k *Kubernetes) CreateUser(name string) error {
	clientset, err := kubernetes.NewForConfig(k.Config)
	if err != nil {
		return err
	}

	// create service account if it doesn't exist for this user
	_, err = clientset.CoreV1().ServiceAccounts("default").Create(context.Background(), &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: normalize(name),
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "infra",
			},
		},
	}, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

func (k *Kubernetes) Grant(user string, role string) error {
	clientset, err := kubernetes.NewForConfig(k.Config)
	if err != nil {
		return err
	}

	name := normalize(user)

	_, err = clientset.RbacV1().ClusterRoleBindings().Create(context.Background(), &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "infra",
			},
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      rbacv1.ServiceAccountKind,
				Name:      name,
				Namespace: "default",
			},
		},
		RoleRef: rbacv1.RoleRef{
			Kind: "ClusterRole",
			Name: role,
		},
	}, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}

	return nil
}
