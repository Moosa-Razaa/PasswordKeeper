const baseUrl: string = "http://localhost:3000";

function BuildUrl(resource: string): string
{
    return `${baseUrl}/${resource}`;
}

export default BuildUrl;