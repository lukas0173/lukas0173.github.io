<script>
    import {Content} from "./Tab/index.js";

    export let projects, activeTabId, imagePath, technologyColors
</script>

{#each projects as projectType, projectTypeIndex}
    {#if projectType.projectList === null}
        <Content {activeTabId}
                 class="flex flex-col gap-10 my-auto items-center"
                 id={projectTypeIndex}>
            <img src="assets/images/Main/Projects/nothing.png" alt="nothing"
                 class="w-28 md:w-32 xl:w-40">
            <h2 class="text-lg md:text-xl xl:text-2xl font-semibold opacity-50">Nothing yet</h2>
        </Content>
    {:else}
        <Content {activeTabId}
                 class="flex flex-col md:grid md:grid-cols-2 gap-8 transition-all"
                 id={projectTypeIndex}>
            {#each projectType.projectList as project, projectIndex}
                <div class="flex flex-col 2xl:flex-row gap-5 bg-Base rounded-xl p-5">
                    <img class="object-cover 2xl:w-[calc(100vw/6)] aspect-[6/3] 2xl:aspect-[4/3] rounded-lg"
                         src={`${imagePath}/${projectType.imageFolder}/${project.imageName}.png`}
                         alt="project"/>
                    <div class="flex h-full flex-col gap-2">
                        <a href={project.link} target="_blank">
                            <h3 class="text-xl lg:text-2xl leading-none font-semibold">{project.projectName}
                                <i class="ri-external-link-fill"> </i>
                            </h3>
                        </a>
                        <p class="text-sm lg:text-base">{project.projectDescription}</p>
                        <div class="mt-2 flex gap-3">
                            {#each project.technologies as technology, technologyIndex}
                                <!--accessTechnologyColors with the root index (projects index), then the project inside the projectList property, and finally, the color property-->
                                <p class={"rounded-lg text-sm text-Base px-2 py-[2px] " + technologyColors[projectTypeIndex].projectList[projectIndex].colors[technologyIndex]}>{technology}</p>
                            {/each}
                        </div>
                    </div>
                </div>
            {/each}
        </Content>
    {/if}
{/each}

