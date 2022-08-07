/**
 * GAA
 * 1.0.0
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "oazapfts/lib/runtime";
import * as QS from "oazapfts/lib/runtime/query";
export const defaults: Oazapfts.RequestOpts = {
    baseUrl: "http://localhost:8080/api",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {
    server1: "http://localhost:8080/api"
};
export type RowsResult = {
    ModelName?: string;
    TableName?: string;
    Fields?: {
        Name?: string;
        Type?: string;
        Size?: number;
    }[];
    Data?: any[];
};
/**
 * Returns all model names
 */
export function getModels(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: string[];
    }>("/models", {
        ...opts
    });
}
/**
 * Returns model rows
 */
export function getModelsByModelRows(model: string, { limit, offset }: {
    limit?: number;
    offset?: number;
} = {}, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: RowsResult;
    }>(`/models/${model}/rows${QS.query(QS.form({
        limit,
        offset
    }))}`, {
        ...opts
    });
}
