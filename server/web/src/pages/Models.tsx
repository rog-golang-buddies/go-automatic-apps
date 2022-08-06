
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import {defaults, getModels} from "../api/api";

defaults.baseUrl = "http://localhost:8080/api/";

export default function Models() {

    const [models, setModels] = useState([] as string[]);

    useEffect(() => {
        getModels().then(res => {
            if (res.status === 200) {
                setModels(res.data);
            } else {
                setModels([]);
                console.log("Error", res);
            }
        });
    }, []);
    
    return <>
        <h1>Models</h1>
        {models.map(model => 
            <div>
                Model: <Link to={"/models/" + model}>Go to {model}</Link>
            </div>
        )}
        {models.length === 0 && <div>No models</div>}

        <Link to={"/"}>Go back</Link>
    </>
}