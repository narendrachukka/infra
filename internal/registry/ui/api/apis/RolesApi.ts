/* tslint:disable */
/* eslint-disable */
/**
 * Infra API
 * Infra REST API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    Role,
    RoleFromJSON,
    RoleToJSON,
} from '../models';

export interface GetRoleRequest {
    id: string;
}

export interface ListRolesRequest {
    name?: string;
    kind?: string;
    destination?: string;
}

/**
 * 
 */
export class RolesApi extends runtime.BaseAPI {

    /**
     * Get role
     */
    async getRoleRaw(requestParameters: GetRoleRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Role>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getRole.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/roles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RoleFromJSON(jsonValue));
    }

    /**
     * Get role
     */
    async getRole(requestParameters: GetRoleRequest, initOverrides?: RequestInit): Promise<Role> {
        const response = await this.getRoleRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List roles
     */
    async listRolesRaw(requestParameters: ListRolesRequest, initOverrides?: RequestInit): Promise<runtime.ApiResponse<Array<Role>>> {
        const queryParameters: any = {};

        if (requestParameters.name !== undefined) {
            queryParameters['name'] = requestParameters.name;
        }

        if (requestParameters.kind !== undefined) {
            queryParameters['kind'] = requestParameters.kind;
        }

        if (requestParameters.destination !== undefined) {
            queryParameters['destination'] = requestParameters.destination;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/roles`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(RoleFromJSON));
    }

    /**
     * List roles
     */
    async listRoles(requestParameters: ListRolesRequest, initOverrides?: RequestInit): Promise<Array<Role>> {
        const response = await this.listRolesRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
