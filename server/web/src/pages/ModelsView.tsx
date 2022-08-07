import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";

import {defaults, getModelsByModelRows, RowsResult} from "../api/api";

defaults.baseUrl = "http://localhost:8080/api/";

export default function ModelView() {

    const params = useParams();
    const modelName = params.modelName || "";

    const [rowsResult, setRowsResult] = useState({} as RowsResult);

    useEffect(() => {
        getModelsByModelRows(modelName, {limit: 10, offset: 0}).then(res => {
            if (res.status === 200) {
                setRowsResult(res.data);
            } else {
                setRowsResult({});
                console.log("Error", res);
            }
        });
    }, []);
    
    return <>
        <h1>Model: {rowsResult.ModelName}</h1>
        <p>Table: {rowsResult.TableName}</p>
        <p>Fields:</p>
        <div>
            {JSON.stringify(rowsResult.Fields)}
        </div>
        <p>Data:</p>
        {rowsResult.Data?.map((row, i) => 
            <div>
                Row {i+1}: {JSON.stringify(row)}
            </div>
        )}
        {rowsResult.Data?.length === 0 && <div>No data.</div>}

        <Link to={"/"}>Go back</Link>
    </>
}