<script lang="ts">
import {UserCreateUrlValidation} from "$lib";

let nullAlert:Boolean = false

let sourceUrlValidation:Boolean = true

function ValidURL(){
    const input = <HTMLInputElement>document.getElementById("source")
    sourceUrlValidation = UserCreateUrlValidation(input.value)
}
</script>


<form method="POST">
    <label class="label">
        <span>Destination</span>
        <input name="destination" class="input" type="text" placeholder="e.g https://www.google.com/"/>
    </label>
    <label class="label input-group input-group-divider grid-cols-[auto_1fr_auto]">
        <span class="inline-block text-lg justify-between  text-center p-2">lat.sh/</span>
        <input id="source" name="source" class="input variant-form-material" class:input-success={sourceUrlValidation} class:input-error={!sourceUrlValidation} on:input={ValidURL} type="text" placeholder="e.g Hello_world"/>
    </label>
    <br>
    <button type="submit" class="btn variant-filled-surface">Create</button>
</form>


{#if nullAlert}
    <aside class="alert variant-ghost">
        <div class="alert-message">
           <h3 class="h3">Something went wrong</h3>
            <p>oops looks like something went wrong, please refresh the page</p>
        </div>
    </aside>
{/if}