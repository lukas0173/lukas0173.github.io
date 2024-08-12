<script>
    let positions = [{
        role: "Front-end developer",
        color: "bg-Teal",
        shadowColor: "shadow-Teal"
    }, {
        role: "Back-end developer", color: "bg-Red",
        shadowColor: "shadow-Red"
    }, {
        role: "Linux administration",
        color: "bg-Green",
        shadowColor: "shadow-Green"
    }, {
        role: "Embedded developer", color: "bg-Lavender",
        shadowColor: "shadow-Lavender"
    }]
</script>

<!--The flex style here is just to center the div inside it since my-auto is not working-->
<section id="About me" class="min-h-screen flex items-center my-10 lg:my-40">
    <div class="flex flex-col lg:grid lg:grid-cols-2 gap-10 md:gap-20">
        <img class="hidden lg:inline place-self-center lg:w-[calc(100vw/3)] xl:w-[calc(100vw/3.5)] 2xl:w-[calc(100vw/4)]" src="assets/images/Main/About_me/illustration.png"
             alt="illustration"/>
        <div class="flex flex-col gap-5">
            <h1 class="leading-none text-6xl font-bold">About <span class="text-Teal">Me</span></h1>
            <p class="text-base md:text-lg xl:text-xl">Concentration is one of the strengths that I'm self-assured. This helps me make use of as
                much time and effort as possible to complete the project ahead of time. Another advantage is that I'm
                not afraid to acquire new technology, especially when understanding the basics.</p>
            <div class="flex gap-2 text-base md:text-lg xl:text-xl leading-none place-items-center">
                <p>I will be your</p>
                <i class="ri-arrow-down-circle-line mt-1"></i>
            </div>
            <div class="place-self-center relative lg:mr-auto grid grid-cols-2 gap-6 mt-5 md:mt-10 lg:my-0">
                {#each positions as position}
                    <div id="position"
                         class={"px-6 py-5 rounded-lg shadow-[0px_5px_100px_-15px_rgba(0,0,0,0.25)] " + position.color + " " + position.shadowColor}>
                        <p class="leading-none text-base md:text-lg xl:text-xl text-center font-medium text-Base">{position.role}</p>
                    </div>
                {/each}
            </div>
        </div>
        <img class="lg:hidden self-center h-[calc(100vh/4)] md:h-[calc(100vh/3)]" src="assets/images/Main/About_me/illustration.png"
             alt="illustration"/>
    </div>
</section>
