import {HOST_ADDRESS} from "$env/static/private"

export async function load({fetch, params}) {
    let returnObject = {};
    let developmentToolsResponse = await fetch(`http://${HOST_ADDRESS}/development-tools`)
    let personalProjectsResponse = await fetch(`http://${HOST_ADDRESS}/personal-projects`)
    let teamProjectsResponse = await fetch(`http://${HOST_ADDRESS}/team-projects`)
    returnObject = {
        developmentTools: await developmentToolsResponse.json(),
        personalProjects: await personalProjectsResponse.json(),
        teamProjects: await teamProjectsResponse.json()
    }
    return returnObject
}
