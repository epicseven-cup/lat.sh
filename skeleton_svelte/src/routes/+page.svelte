<script lang="ts">
    import {UserCreateUrlValidation} from "$lib";
    import {fade} from "svelte/transition";


    let errorAlert: Boolean = false
    let sourceUrlValidation: Boolean = true
    const message: string = "Something went wrong"


    let source: string = ""
    let destination: string = ""
    function ValidURL() {
        sourceUrlValidation = UserCreateUrlValidation(source)
    }

    let successAlert: boolean = false;
    async function handleSubmit(){
        let data = {
            "source": source,
            "destination": destination
        }
        let respond:Response = await fetch("api/generate", {
            method: 'POST',
            body: JSON.stringify(data)
        })

        if (respond.ok){
            successAlert= true
            setTimeout(() => successAlert = false, 5000)
        } else {
            errorAlert = true
            setTimeout(() => errorAlert = false, 5000)
        }
    }


</script>


<div class="text-center">
    <div class="inline-block text-left">
        <pre class="bg-gradient-to-br from-red-500 to-yellow-500 bg-clip-text text-transparent box-decoration-clone">
             _       _         _
            | | __ _| |_   ___| |__
            | |/ _` | __| / __| '_ \
            | | (_| | |_ _\__ \ | | |
            |_|\__,_|\__(_)___/_| |_|
        </pre>
    </div>
</div>
<br>
<div class="card p-4">
    <h1 class="card-header text-center h2">
        <span>Create new Links</span>
    </h1>
        <form method="POST"  on:submit|preventDefault={handleSubmit}>


        <label class="label">
            <span>Destination</span>
            <input name="destination" bind:value={destination} class="input" type="text" placeholder="e.g https://www.google.com/"/>
        </label>

        <label class="label input-group input-group-divider grid-cols-[auto_1fr_auto]">
            <span class="inline-block text-lg justify-between text-center p-2">lat.sh/</span>
            <input id="source" name="source" class="input variant-form-material"
                   class:input-success={sourceUrlValidation} class:input-error={!sourceUrlValidation}
                   on:input={ValidURL} bind:value={source} type="text" placeholder="e.g Hello_world"/>
        </label>

        <br>
        <button type="submit" class="btn variant-filled-surface">Create</button>
    </form>


</div>


{#if errorAlert}
    <aside class="alert variant-ringed"  transition:fade={{ duration: 100 }}>
        <div class="alert-message">
            <p>{message}</p>
        </div>
    </aside>
{/if}

{#if successAlert}
    <aside class="alert variant-ringed"  transition:fade={{ duration: 100 }}>
            <p>Url is created</p>
    </aside>
{/if}
