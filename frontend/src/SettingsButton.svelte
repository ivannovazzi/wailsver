<script>
  import Tooltip from './Tooltip.svelte';
  export let onClick;
</script>

<div class="settings-button-container">
  <Tooltip text="Settings" position="right">
    <button class="settings-button" on:click={onClick}>
      <!-- Gear Icon SVG -->
      <svg
        width="24"
        height="24"
        fill="#1890ff"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M12 1.75a10.249 10.249 0 100 20.5 10.249 10.249 0 000-20.5zm0 18.5a8.251 8.251 0 110-16.5 8.251 8.251 0 010 16.5zm1.5-13.5h-3v3h3v-3zm0 4.5h-3v6h3v-6z"
        />
      </svg>
    </button>
  </Tooltip>
</div>

<style>
  .settings-button-container {
    top: 20px;
    left: 20px;
  }

  .settings-button {
    background: none;
    border: none;
    cursor: pointer;
  }
</style>