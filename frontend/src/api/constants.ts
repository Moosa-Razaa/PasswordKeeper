const baseUrl: string = "http://localhost:8080";

function BuildUrl(resource: string): string
{
    return `${baseUrl}/${resource}`;
}

export default BuildUrl;