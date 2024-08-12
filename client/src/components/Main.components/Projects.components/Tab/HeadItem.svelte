<script>
    export let id;
</script>

<li class="w-full">
    <button
            on:click
            class={$$restProps.class || ""}
            {id}
            type="button"
            role="tab"
    >
        <slot/>
    </button>
</li>
