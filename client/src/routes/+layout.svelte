<script>

    import "../app.css";
    import Footer from "/src/components/Footer.svelte";
    import Header from "/src/components/Header.svelte";

    import {setContext} from "svelte";
    import {writable} from "svelte/store";

    let scrollY = writable(null);
    let innerHeight = 0,
        innerWidth = 0;

    const goTop = () => {
        document.body.scrollIntoView();
    };

    //Passing scrollY to child elements
    setContext("scroll", {scrollY});
</script>

<div class="relative flex flex-col w-full min-h-svh">
    <Header/>
    <slot/>
    <Footer/>
</div>

<svelte:window bind:scrollY={$scrollY} bind:innerHeight bind:innerWidth/>
