<script>
    import Card from "./Card.svelte"
    import Modal from "./Modal.svelte"
    import { Modals, closeModal, openModal } from "svelte-modals"

    export let shurl
    let showCard = true
    async function update(data) {
        const json = {
            redirect: data.redirect,
            shurl: data.shurl,
            random: data.random,
            id: shurl.id
        }

        await fetch("http://localhost:3000/shurl", {
            method: "PATCH",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res)
        })
    }

    function handleOpen(shurl) {
        openModal(Modal, {
            title: "Update Short URL Link",
            send: update,
            shurl: shurl.shurl,
            redirect: shurl.redirect,
            random: shurl.random
        })
    }

    async function deleteShurl() {
        if (confirm("Are you sure you wish to delete this short url link (" + shurl.shurl + ")?")) {
            await fetch("http://localhost:3000/shurl/" + shurl.id, {
                method: "DELETE"
            }).then(response => {
                showCard = false
                console.log(response)
            })
        }
    }

    function copyShurl() {
        const shortedUrl = `http://localhost:3000/r/${shurl.shurl}`;
        navigator.clipboard.writeText(shortedUrl).then(() => {
            alert("Copied to clipboard");
        });
    }

</script>
{#if showCard }
    <Card>
        <p>Short url:<a href="http://localhost:3000/r/{shurl.shurl}">" http://localhost:3000/r/{shurl.shurl}"</a></p>
        <p>Long url: {shurl.redirect}</p>
        <p>Clicked: {shurl.clicked}</p>
        <button class="update" on:click={ handleOpen(shurl) }>Update</button>
        <button class="delete" on:click={deleteShurl}>Delete</button>
        <button class="copy" on:click={copyShurl}>Copy</button>
    </Card>
{/if}

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }
    .update {
        background-color: yellowgreen;
    }
    .delete {
        background-color: red;
    }
    .copy {
        background-color: blue;
    }
    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: rgb(255, 255, 255)
    }

</style>