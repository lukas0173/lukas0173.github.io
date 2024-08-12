import {API_URL} from "$env/static/private"

export async function load({fetch, params}) {
    let returnObject = {};
    let developmentToolsResponse = await fetch(`${API_URL}/development-tools`)
    let personalProjectsResponse = await fetch(`${API_URL}/personal-projects`)
    let teamProjectsResponse = await fetch(`${API_URL}/team-projects`)
    returnObject = {
        developmentTools: await developmentToolsResponse.json(),
        personalProjects: await personalProjectsResponse.json(),
        teamProjects: await teamProjectsResponse.json()
    }
    return returnObject
}