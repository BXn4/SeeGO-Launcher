<script lang="ts">
    import logo from "../public/images/seego_nobg.svg";
</script>

<main>
    <div id="splash">
        <img id="splash-logo" src={logo} alt="SeeGO Launcher logo" />
        <h1 id="splash-title">SeeGO Launcher</h1>
        <p id="splash-alt">Betöltés</p>

        <div id="loading"></div>
    </div>
</main>

<style>
    :global(body, html) {
        color: #ffffff;
        margin: 0;
        padding: 0;
        width: 100%;
        height: 100%;
        overflow: hidden;
        font-family:
            system-ui,
            -apple-system,
            Segoe UI,
            Roboto,
            sans-serif;
    }

    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;

        background: linear-gradient(
            -30deg,
            #1b2636,
            #1b2534,
            #1a2433,
            #1a2331,
            #1a222f,
            #1a212e,
            #19202c,
            #191f2a,
            #191e29,
            #181e27,
            #181d25,
            #171c24,
            #171b22,
            #161a21,
            #16191f
        );
        background-size: 400% 400%;
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

    #splash-alt {
        margin: 0;
        opacity: 0.7;
        font-size: 14px;
    }

    #loading {
        margin-top: 16px;
        width: 28px;
        height: 28px;
        border: 3px solid rgba(255, 255, 255, 0.2);
        border-top: 3px solid white;
        border-radius: 50%;
        animation: spin 1.5s linear infinite;
    }

    @keyframes spin {
        to {
            transform: rotate(360deg);
        }
    }

    @keyframes popIn {
        from {
            transform: scale(0.7);
            opacity: 0;
        }
        to {
            transform: scale(1);
            opacity: 1;
        }
    }

    @keyframes fadeUp {
        from {
            transform: translateY(10px);
            opacity: 0;
        }
        to {
            transform: translateY(0);
            opacity: 1;
        }
    }
</style>
