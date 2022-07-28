import { useState } from "react";
import { Link } from "react-router-dom";




export default function Models() {

    const [models, setModels] = useState([]);

    // TODO: call gRPC service to get models

    return <>
        <h1>Models</h1>
        {models.map(model => <div>{model}</div>)}
        {models.length === 0 && <div>No models</div>}

        <Link to={"/"}>Go back</Link>
    </>
}