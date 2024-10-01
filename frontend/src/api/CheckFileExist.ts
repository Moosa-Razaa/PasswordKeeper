import axios from "axios";
import BuildUrl from "./constants";

async function CheckFileExist(filePath: string): Promise<string> 
{
    try 
    {
        const url: string = BuildUrl("/check/" + filePath);
        const response = await axios.get(url);
        
        if (response.status !== 200) 
        {
            return response.data;
        }
    } 
    catch (error) 
    {
        return error.message;
    }
    finally 
    {
        return "";
    }
}

export default CheckFileExist;