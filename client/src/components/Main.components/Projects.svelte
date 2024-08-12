<script>
    import {getContext, onMount} from "svelte";
    import * as Tab from "./Projects.components/Tab"
    import TabHead from "./Projects.components/TabHead.svelte";
    import TabContent from "./Projects.components/TabContent.svelte";

    // get the personal and team project data from context
    const personalProjects = getContext("personalProjects")
    const teamProjects = getContext("teamProjects")
    const projects = [
        {
            name: "Personal works",
            imageFolder: "personal_work",
            projectList: personalProjects
        },
        {
            name: "Team projects",
            imageFolder: "team_projects",
            projectList: teamProjects
        }
    ]


    let activeTabId = 0;

    const imagePath = "assets/images/Main/Projects";

    // Technologies background color randomizer
    const colors = ["bg-Red", "bg-Green", "bg-Blue", "bg-Pink", "bg-Yellow", "bg-Teal", "bg-Peach"];
    // Return the exact structure of project object, only with the colors property
    const technologyColors = projects.map((projectType) => {
        return {
            name: projectType,
            projectList: (projectType.projectList === null) ? null : projectType.projectList.map((project) => {
                // This is for resetting the temp color array for each project
                let tempColors = [...colors];
                return {
                    projectName: project.projectName,
                    colors: project.technologies.map(() => {
                        //Reset the temp colors value after it's empty, cannot reset it outside the map method since it have to wait for the whole proj
                        const randomNumber = Math.floor(Math.random() * tempColors.length);
                        const returnedColor = tempColors[randomNumber];
                        //Prevent repeating color
                        tempColors.splice(randomNumber, 1);
                        return returnedColor;
                    })
                }
            }),
        }
    })


    const handleClick = (tabValue) => {
        activeTabId = tabValue;
    };

    //Animated indicator

    // The TailwindCss "transition" does not know how to transit the background horizontally
    // between component if you only specify the background color, thus a distinct component
    // is a proper solution for this task.

    let tabs, indicator;

    const updateTabIndicator = (tabs, indicator) => {
        indicator.style.width = tabs[0].getBoundingClientRect().width + "px";
        indicator.style.height = tabs[0].getBoundingClientRect().height + "px";
        indicator.style.left = tabs[0].getBoundingClientRect().left - tabs[0].parentElement.getBoundingClientRect().left + "px"

        tabs.forEach((tab) => {
            tab.addEventListener("click", () => {
                indicator.style.left = tab.getBoundingClientRect().left - tabs[0].parentElement.getBoundingClientRect().left + "px"
            })
        })
    }
    // To check on the resize event
    let previousHeight;

    onMount(() => {
            // Initialize the height value
            previousHeight = window.innerHeight;

            tabs = document.querySelectorAll(".tab");
            indicator = document.querySelector("#indicator");
            updateTabIndicator(tabs, indicator);
        }
    )

</script>

<!--Resize the tab indicator on width change-->
<svelte:window on:resize={() => {
    if (window.innerHeight !== previousHeight) {
        previousHeight = window.innerHeight
        return
    }
    updateTabIndicator(tabs, indicator)
    }
}/>

<section class="flex flex-col gap-10 pt-20 md:pt-32 lg:pt-40" id="Projects">
    <h2
            class="text-center font-semibold text-2xl md:text-4xl lg:text-5xl"
            id="Heading"
    >
        Curious to see my <span class="text-Teal">creations</span>?
    </h2>
    <Tab.Root class="flex flex-col min-h-[calc(100svh/2)] gap-10" id="Body">
        <TabHead {activeTabId} {handleClick} {projects}/>
        <TabContent {activeTabId} {imagePath} {projects} {technologyColors}/>
    </Tab.Root>
</section>
