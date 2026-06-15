<script lang="ts">
    import logo from "../public/images/seego_nobg.svg";
</script>

<main>
    <div id="splash">
        <img id="splash-logo" src={logo} alt="SeeGO Launcher logo" />
        <h1 class="text splash-title">SeeGO Launcher</h1>
        <span class="loader"></span>
    </div>
</main>

<style>
    @import "../public/styles/views/splash.css";

    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    :global(body, html) {
        margin: 0;
        padding: 0;
    }
</style>
