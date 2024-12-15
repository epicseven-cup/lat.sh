<script lang="ts">
    import {UserCreateUrlValidation} from "$lib";


</script>


<form method="POST">
    <label class="label">
        <span>Destination</span>
        <input name="destination" class="input" type="text" placeholder="e.g https://www.google.com/"/>
    </label>
    <label class="label input-group input-group-divider grid-cols-[auto_1fr_auto]">
        <span class="inline-block text-lg justify-between  text-center p-2">lat.sh/</span>
        <input name="source" class="input variant-form-material" type="text" placeholder="e.g Hello_world"/>
    </label>
    <br>
    <button type="submit" class="btn variant-filled-surface">Create</button>
</form>