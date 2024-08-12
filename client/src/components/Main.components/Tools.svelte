<script>
    import {getContext} from "svelte";

    // Get the developmmentTools object
    const developmentTools = getContext("developmentTools")

    // Dynamic styles for column span and text color
    const toolSectionStyles = developmentTools.map((tool) => {
        return {span: "col-span-" + tool.style.span, text_color: "text-" + tool.style.textColor}
    })

    const iconsPath = "assets/images/Main/Tools";
</script>

<section
        id="Tools"
        class="flex flex-col gap-5 md:gap-10 pt-10 lg:pt-32"
>
    <div id="Heading" class="flex flex-col gap-4 lg:gap-5">
        <div
                class="bg-Teal animate-[spin_2.5s_linear_infinite] mx-auto rounded-full p-3 lg:p-4 w-fit text-Base leading-none text-lg lg:text-3xl lg:leading-none"
        >
            <i class="ri-settings-5-fill"></i>
        </div>
        <h2 class="font-semibold text-center text-3xl md:text-4xl lg:text-5xl">
            Favourite development tools
        </h2>
    </div>

    <div id="Body" class="flex flex-col lg:grid lg:grid-cols-10 gap-7 lg:gap-5">
        {#each developmentTools as tool, index}
            <!--Can use the styles index directly because the map method in the script
            mapped the styles at the same index as the tool's field-->
            <div
                    class={"p-7 flex flex-col gap-5 bg-Base rounded-lg " + toolSectionStyles[index].span}
            >
                <h3
                        class={"leading-none font-semibold text-2xl " + toolSectionStyles[index].text_color}
                >
                    {tool.field}
                </h3>
                <p class="text-base md:text-lg">{tool.description}</p>
                <div
                        class={"self-center mt-auto flex items-center max-w-full justify-between flex-wrap lg:flex-nowrap " +
            (tool.icons.name.length <= 4
              ? tool.icons.name.length <= 2
                ? "w-1/2 lg:w-2/3 xl:w-3/5"
                : "w-full sm:w-2/3 md:w-3/5 lg:w-full 2xl:w-10/12"
              : "gap-y-2 w-full sm:w-2/3 md:w-3/5 lg:w-full 2xl:w-10/12")}
                >
                    {#each tool.icons.name as iconName (iconName)}
                        <img
                                class={iconName === "IntelFPGA" ? "h-6 md:h-8" : "h-10 md:h-14"}
                                src={`${iconsPath}/${tool.icons.path}/${iconName}.png`}
                                alt={iconName}
                        />
                        <!--This is for break the icons in 2 rows in mobile view-->
                        {#if iconName === "JavaScript"}
                            <div class="basis-full lg:hidden"></div>
                        {/if}
                    {/each}
                </div>
            </div>
        {/each}
    </div>
</section>
