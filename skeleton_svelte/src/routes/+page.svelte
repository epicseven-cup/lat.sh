<script lang="ts">
    import {UserCreateUrlValidation} from "$lib";
    import {fade} from "svelte/transition";


    // Deteremine if there should be an error or success
    let errorAlert: Boolean = false
    let successAlert: boolean = false;
    let sourceUrlValidation: Boolean = true


    const fadeoutTime: number = 5000

    let source: string = ""
    let destination: string = ""
    function ValidURL() {
        sourceUrlValidation = UserCreateUrlValidation(source)
    }

    // This can be done in svelete with some dynamic elements, though for the time being raw js it is
    function addAlert(message: string, status: bool){
    	let parent = document.getElementById("alertSelection")
	let child = document.createElement("aside")
	child.classList.add("alert")
	if (status){
		child.classList.add("variant-filled-success")
	} else {
		child.classList.add("variant-filled-error")
	}
	// Removing the element once the timer hit in
	setTimeout(() => child.remove(), fadeoutTime)
	let messageElement = document.createElement("div")
	messageElement.classList.add("alert-message")
	messageElement.innerHTML = message

	child.appendChild(messageElement)
	parent.prepend(child)
    }

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
            addAlert("URL Created", true)
        } else {
            addAlert("Something Went Wrong")
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
            <span class="inline-block text justify-between text-center p-2">lat.sh/</span>
            <input id="source" name="source" class="input variant-form-material"
                   class:input-success={sourceUrlValidation} class:input-error={!sourceUrlValidation}
                   on:input={ValidURL} bind:value={source} type="text" placeholder="e.g Hello_world"/>
        </label>

        <br>
        <button type="submit" class="btn variant-filled-surface">Create</button>
    </form>
</div>

<div id="alertSelection"></div>
