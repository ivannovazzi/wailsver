<!-- SettingsModal.svelte -->
<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { main } from '../wailsjs/go/models';

  export let config: main.Config

  const dispatch = createEventDispatcher();

  function saveSettings() {
    dispatch('save', config);
    close();
  }

  function isConfigValid() {
    return config.github_token && config.owner && config.repo && config.table_name;
  }

  function close() {
    dispatch('close');
  }
</script>

<div class="modal">
  <div class="modal-content">
    <span class="close" on:click={close} role="presentation">&times;</span>
    <h2>Settings</h2>
    <div class="input-box">
      <input bind:value={config.github_token} class="input" type="text" placeholder="GitHub Token" />
      <input bind:value={config.owner} class="input" type="text" placeholder="Repo Owner" />
      <input bind:value={config.repo} class="input" type="text" placeholder="Repo Name" />
      <input bind:value={config.table_name} class="input" type="text" placeholder="Table Name" />
      <button class="btn" on:click={saveSettings} disabled={!isConfigValid()}>Save</button>
    </div>
  </div>
</div>

<style>
  .modal {
    display: block;
    position: fixed;
    z-index: 1;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.4);
    -webkit-backdrop-filter: blur(3px);
    backdrop-filter: blur(3px);
    box-shadow: 0px 0px 50px rgba(0, 0, 0, 0.4);
  }

  .modal-content {
    background-color: #50505050;
    -webkit-backdrop-filter: blur(10px);
    backdrop-filter: blur(10px);
    margin: 10% auto;
    padding: 20px;
    border-radius: 8px;
    width: 80%;
    max-width: 500px;
    position: relative;
    box-shadow: 0px 0px 50px rgba(0, 0, 0, 0.4);
    /* make a fade-in effect when appears, use a professional fading curve */
    animation: fadeIn 0.3s cubic-bezier(0.39, 0.575, 0.565, 1) both;
  }


  .close {
    color: #aaa;
    position: absolute;
    right: 20px;
    top: 15px;
    font-size: 28px;
    cursor: pointer;
  }

  .close:hover {
    color: black;
  }

  .input-box {
    display: flex;
    flex-direction: column;
    margin-top: 20px;
  }

  .input-box .input {
    flex: 1 1 auto;
    height: 45px;
    margin-bottom: 15px;
    padding: 0 15px;
    border: none;
    border-radius: 8px;
    background-color: rgba(120, 120, 120, 0.8);
    font-size: 1rem;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);    
  }

  .input-box .btn {
    width: 100%;
    height: 45px;
    background-color: #1890ff;
    color: #fff;
    border: none;
    border-radius: 8px;
    font-size: 1rem;
    cursor: pointer;
  }

  .input-box .btn:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  /* create the fadeIn animation with little scaling, opacity, and movement from bottom  */
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: scale(0.8) translateY(30px);
    }
    to {
      opacity: 1;
      transform: translateY(0px);
    }
  }
</style>