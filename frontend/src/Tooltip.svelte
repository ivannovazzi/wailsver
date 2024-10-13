<script>
  export let text = 'Tooltip';
  export let position = 'top'; // options: top, right, bottom, left
  export let backgroundColor = 'rgba(0, 0, 0, 0.6)'; // semi-transparent dark background
  export let textColor = '#fff';
  export let borderRadius = '6px';
  export let padding = '5px 10px';
</script>

<div class="tooltip-container">
  <slot />
  <div
    class="tooltip"
    data-position={position}
    style="
      background-color: {backgroundColor};
      color: {textColor};
      border-radius: {borderRadius};
      padding: {padding};
    "
  >
    {text}
  </div>
</div>

<style lang="scss">
  .tooltip-container {
    position: relative;
    display: inline-block;
  }

  .tooltip {
    position: absolute;
    z-index: 1;
    visibility: hidden;
    opacity: 0;
    transition: opacity 0.3s;
    white-space: nowrap;
    text-align: center;
    -webkit-backdrop-filter: blur(5px); // Apply blur effect
    backdrop-filter: blur(5px); // Apply blur effect
  }

  .tooltip-container:hover .tooltip {
    visibility: visible;
    opacity: 1;
  }

  .tooltip[data-position='top'] {
    bottom: 100%;
    left: 50%;
    transform: translateX(-50%);
    margin-bottom: 8px;
  }

  .tooltip[data-position='right'] {
    top: 50%;
    left: 100%;
    transform: translateY(-50%);
    margin-left: 8px;
  }

  .tooltip[data-position='bottom'] {
    top: 100%;
    left: 50%;
    transform: translateX(-50%);
    margin-top: 8px;
  }

  .tooltip[data-position='left'] {
    top: 50%;
    right: 100%;
    transform: translateY(-50%);
    margin-right: 8px;
  }
</style>