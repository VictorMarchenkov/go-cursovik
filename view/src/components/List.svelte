<script>
    import Item from "./Item.svelte"
    import { onMount } from "svelte"

    let shurls = []

    onMount(async () => {
        const res = await fetch("http://localhost:3000/shurl")
        shurls = await res.json()
    })
</script>

{#each shurls as shurl}
    <Item shurl={shurl} />
{/each}
