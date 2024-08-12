<script>
    import {fade} from "svelte/transition"
    export let activeTabId;
    export let id;
</script>

{#if activeTabId === id}
    <div in:fade={{duration: 200}} class={$$restProps.class || ""} {id} role="tabpanel">
        <slot class={$$restProps.class || ""}></slot>
    </div>
{/if}
