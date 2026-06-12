<script lang="ts">
    import logo from "../public/images/seego_nobg.svg";
    import { onMount } from "svelte";
    import {
        Localization,
        Config,
    } from "../../bindings/seegolauncher/internal/services";
</script>

<main>
    <div id="splash">
        <img id="splash-logo" src={logo} alt="SeeGO Launcher logo" />
        <h1 id="splash-title">SeeGO Launcher</h1>
        <div id="loading"></div>
    </div>
</main>

<style>
    :global(body, html) {
        color: var(--text);
        margin: 0;
        padding: 0;
        width: 100%;
        height: 100%;
        overflow: hidden;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }

    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    #splash {
        text-align: center;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 14px;
    }

    #splash-logo {
        width: 120px;
        height: auto;
    }

    #splash-title {
        font-size: 28px;
        margin: 0;
        font-weight: 600;
        letter-spacing: 0.5px;
    }

    #loading {
        margin-top: 16px;
        width: 28px;
        height: 28px;
        border: 3px solid var(--comment);
        border-top: 3px solid white;
        border-radius: 50%;
        animation: spin 1.5s linear infinite;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }
</style>
