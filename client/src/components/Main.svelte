<script>
    import Introduction from "./Main.components/Introduction.svelte";
    import Projects from "./Main.components/Projects.svelte";
    import Tools from "./Main.components/Tools.svelte";
    import AboutMe from "./Main.components/AboutMe.svelte";
</script>

<main class="flex flex-col px-10 lg:px-20 xl:px-[120px] min-h-svh">
    <Introduction/>
    <Tools/>
    <Projects/>
    <AboutMe/>
</main>
