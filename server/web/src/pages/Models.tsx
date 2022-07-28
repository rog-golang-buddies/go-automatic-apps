import { useState } from "react";




export default function Models() {

    const [models, setModels] = useState([]);

    // TODO: call gRPC service to get models

    return <>
        <h1>Models</h1>
        {models.map(model => <div>{model}</div>)}
        {models.length === 0 && <div>No models</div>}

    </>
}