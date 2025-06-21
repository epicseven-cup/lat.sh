<script lang="ts">
	import { UserCreateUrlValidation } from '$lib';
	import { initializeStores, Toast } from '@skeletonlabs/skeleton';
	import { getToastStore } from '@skeletonlabs/skeleton';

	initializeStores();

	// Deteremine if there should be an error or success
	let sourceUrlValidation: Boolean = true;

	const fadeoutTime: number = 10000;


	const t: ToastSettings = {
		message: 'Url copied to clipboard',
	};
	const toastStore = getToastStore();



	let source: string = '';
	let destination: string = '';
	//TODO: This need to be addressed. It should fail request on client side before sending it
	function ValidURL() {
		sourceUrlValidation = UserCreateUrlValidation(source);
	}

	function copyToClipboard(){
		var copyText = document.getElementById("source") as HTMLInputElement;
  		// Select the text field
  		copyText?.select();
  		copyText?.setSelectionRange(0, 99999); // For mobile devices

  		navigator.clipboard.writeText("https://lat.sh/" + copyText.value);
		toastStore.trigger(t)
	}


	// This can be done in svelete with some dynamic elements, though for the time being raw js it is
	function addAlert(message: string, status: Boolean) {
		let parent = document.getElementById('alertSelection');
		let child = document.createElement('aside');
		child.classList.add('alert');

		let messageElement = document.createElement('div');
		messageElement.classList.add('alert-message');
		messageElement.innerHTML = message;

		child.appendChild(messageElement);
		// Removing the element once the timer hit in
		setTimeout(() => child.remove(), fadeoutTime);
		if (status) {
			child.classList.add('variant-filled-success');

			let copyButton = document.createElement('button')
			copyButton.innerHTML = "Copy"
			copyButton.classList.add("alert-actions")
			copyButton.classList.add("btn")
			copyButton.style.border = "normal solid #000000"
			copyButton.onclick = copyToClipboard
			child.appendChild(copyButton)
		} else {
			child.classList.add('variant-filled-error');
		}
		parent?.prepend(child);
	}




	async function handleSubmit() {
		let data = {
			source: source,
			destination: destination
		};
		let respond: Response = await fetch('api/generate', {
			method: 'POST',
			body: JSON.stringify(data)
		});

		if (respond.ok) {
			//TODO: Add popup alert and allow users to copy their url
			addAlert('URL Created', true);
		} else {
			addAlert('Something Went Wrong', false);
		}
	}
</script>

<div class="text-center">
	<div class="inline-block text-left">
		<pre
			class="bg-gradient-to-br from-red-500 to-yellow-500 bg-clip-text text-transparent box-decoration-clone">
             _       _         _
            | | __ _| |_   ___| |__
            | |/ _` | __| / __| '_ \
            | | (_| | |_ _\__ \ | | |
            |_|\__,_|\__(_)___/_| |_|
        </pre>
	</div>
</div>
<br />
<div class="card p-4">
	<h1 class="card-header text-center h2">
		<span>Create new Links</span>
	</h1>
	<form method="POST" on:submit|preventDefault={handleSubmit}>
		<label class="label">
			<span>Destination</span>
			<input
				name="destination"
				bind:value={destination}
				class="input"
				type="text"
				placeholder="e.g https://www.google.com/"
			/>
		</label>

		<label class="label input-group input-group-divider grid-cols-[auto_1fr_auto]">
			<span class="inline-block text justify-between text-center p-2">lat.sh/</span>
			<input
				id="source"
				name="source"
				class="input variant-form-material"
				class:input-success={sourceUrlValidation}
				class:input-error={!sourceUrlValidation}
				on:input={ValidURL}
				bind:value={source}
				type="text"
				placeholder="e.g Hello_world"
			/>
		</label>

		<br />
		<button type="submit" class="btn variant-filled-surface">Create</button>
	</form>
</div>

<div id="alertSelection"></div>

<Toast />
