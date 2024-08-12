<script>
    import {setContext} from "svelte";

    export let data
    import Main from "/src/components/Main.svelte"

    setContext("developmentTools", data.developmentTools)
    setContext("personalProjects", data.personalProjects)
    setContext("teamProjects", data.teamProjects)
</script>

<Main/>
