<script>
    import {getContext} from "svelte";

    //Update the _scrollY value
    let scrollY = getContext("scroll").scrollY;
    $: _scrollY = $scrollY;


  let navigations = [
    { name: "Introduction", link: "#Introduction" },
    { name: "Tools", link: "#Tools" },
    { name: "Projects", link: "#Projects" },
    { name: "About me", link: "#About me" },
  ];

</script>

<header
        class={"sticky z-10 top-0 px-5 md:px-10 flex flex-row items-center justify-between border-b-[1px] border-solid h-20 lg:h-24 " +
    (_scrollY > 0
      ? "bg-Crust border-Surface2 border-opacity-50"
      : "bg-transparent border-transparent")}
>
    <div class="flex gap-8">
        <h1 class="font-bold text-4xl"><span class="text-Teal">ciel</span>aton</h1>
        <div class="lg:flex gap-8 items-center hidden">
            {#each navigations as element}
                <a href={element.link} class="duration-150 hover:text-Teal">
                    <p>{element.name}</p>
                </a>
            {/each}
        </div>
    </div>
    <div class="flex gap-8 items-center">
        <a
                href="https://github.com/cielaton"
                target="_blank"
                class="hidden sm:inline"
        >
            <i class="ri-github-fill text-4xl"></i>
        </a>
        <a
                href="https://www.linkedin.com/in/võ-tuấn-kiệt/"
                target="_blank"
                class="hidden sm:inline"
        >
            <i class="ri-linkedin-box-fill text-4xl"></i>
        </a>
        <a
                href="mailto:vtuankiet17@gmail.com"
                target="_blank"
                class="relative px-5 py-2 group rounded-lg bg-white overflow-hidden"
        >
            <h4 class="relative text-Base font-medium z-[1]">Get in touch</h4>
            <div
                    class="absolute top-0 right-full w-full h-full bg-Text group-hover:translate-x-full z-0 duration-200"
            ></div>
        </a>
    </div>
</header>
