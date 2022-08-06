import { useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";

import {defaults, getModelsByModelRows} from "../api/api";

defaults.baseUrl = "http://localhost:8080/api/";

export default function ModelView() {

    const params = useParams();
    const modelName = params.modelName || "";

    const [data, setData] = useState([] as any[]);

    useEffect(() => {
        getModelsByModelRows(modelName, {limit: 10, offset: 0}).then(res => {
            if (res.status === 200) {
                setData(res.data);
            } else {
                setData([]);
                console.log("Error", res);
            }
        });
    }, []);
    
    return <>
        <h1>Model: {modelName}</h1>
        {data.map((row, i) => 
            <div>
                Row {i+1}: {JSON.stringify(row)}
            </div>
        )}
        {data.length === 0 && <div>No data.</div>}

        <Link to={"/"}>Go back</Link>
    </>
}