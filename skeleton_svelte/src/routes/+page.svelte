<script lang="ts">
    import {UserCreateUrlValidation} from "$lib";

    let nullAlert: Boolean = false
    let sourceUrlValidation: Boolean = true

    function ValidURL() {
        const input = <HTMLInputElement>document.getElementById("source")
        sourceUrlValidation = UserCreateUrlValidation(input.value)
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
        <form method="POST">
        <label class="label">
            <span>Destination</span>
            <input name="destination" class="input" type="text" placeholder="e.g https://www.google.com/"/>
        </label>
        <label class="label input-group input-group-divider grid-cols-[auto_1fr_auto]">
            <span class="inline-block text-lg justify-between  text-center p-2">lat.sh/</span>
            <input id="source" name="source" class="input variant-form-material"
                   class:input-success={sourceUrlValidation} class:input-error={!sourceUrlValidation}
                   on:input={ValidURL} type="text" placeholder="e.g Hello_world"/>
        </label>
        <br>
        <button type="submit" class="btn variant-filled-surface">Create</button>
    </form>
</div>


{#if nullAlert}
    <aside class="alert variant-ghost">
        <div class="alert-message">
            <h3 class="h3">Something went wrong</h3>
            <p>oops looks like something went wrong, please refresh the page</p>
        </div>
    </aside>
{/if}