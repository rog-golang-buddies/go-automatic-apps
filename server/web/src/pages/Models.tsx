import { useState } from "react";




export default function Models() {

    const [models, setModels] = useState([]);

    return <>
        <h1>Models</h1>
        {models.map(model => <div>{model}</div>)}
        {models.length === 0 && <div>No models</div>}

    </>
}