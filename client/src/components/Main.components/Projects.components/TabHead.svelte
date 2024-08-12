<script>
    import * as Tab from "./Tab"

    export let projects, activeTabId, handleClick
</script>

<Tab.Head
        class="relative w-full md:w-4/6 lg:w-3/6 xl:w-2/6 flex items-center mx-auto px-2 py-2 bg-Surface0 rounded-lg">
    <div class="absolute transition-all duration-200 my-auto mx-2 top-0 left-0 bottom-0 rounded-lg bg-Mantle"
         id="indicator"/>
    {#each projects as projectType, index}
        <Tab.HeadItem
                class={"tab relative w-full flex gap-3 py-2 justify-center rounded-lg text-sm md:text-base " + (activeTabId === index ? "text-Teal" : "")}
                id={index}
                on:click={() => handleClick(index)}>
            {#if projectType.name === "Personal works"}
                <i class="ri-user-3-fill"></i>
            {:else if projectType.name === "Team projects"}
                <i class="ri-team-fill"></i>
            {/if}
            <p class="">{projectType.name}</p>
        </Tab.HeadItem
        >
    {/each}
</Tab.Head>
