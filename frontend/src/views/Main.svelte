<script lang="ts">
    import Titlebar from "./partials/Titlebar.svelte";
    import Navbar from "./partials/Navbar.svelte";

    let currentPlayers = 830;
    let maxPlayers = 830;
</script>

<main>
    <Titlebar />
    <Navbar />

    <div id="home-view">
        <div class="feed-layout">
            <header class="hero-card">
                <span class="hero-overlay"></span>
                <div class="hero-content">
                    <span class="badge">Friss</span>
                    <h2>Megnyitottuk az Adminsegéd Tagfelvételt!</h2>
                    <p>
                        Jelentkezési lapok beadásának határideje: 2026.06.12-ig.
                    </p>
                    <button class="hero-news-button" id="hero-news-read-latest"
                        >Elolvasom</button
                    >
                </div>
            </header>

            <div class="items">
                <h3 class="items-title">SeeRPG Club tagság</h3>
                <div class="items-grid">
                    <div class="item-square-card">
                        <div class="item-header">
                            <span class="item-tag">7 nap</span>
                        </div>
                        <div class="item-image">
                            <img
                                src="https://dunb17ur4ymx4.cloudfront.net/packages/images/3f0789ad03f2dad50ebaa1c60b9c8dcfd259dd1a.png"
                                alt="Smaragd"
                            />
                        </div>
                        <div class="item-name">
                            <h4>Smaragd</h4>
                            <span class="item-price">12.19 EUR</span>
                        </div>
                    </div>

                    <div class="item-square-card">
                        <div class="item-header">
                            <span class="item-tag">7 nap</span>
                        </div>
                        <div class="item-image">
                            <img
                                src="https://dunb17ur4ymx4.cloudfront.net/packages/images/7118e68e259c38fd947e9660fc930ae5efa29017.png"
                                alt="Gyémánt"
                            />
                        </div>
                        <div class="item-name">
                            <h4>Gyémánt</h4>
                            <span class="item-price">6.99 EUR</span>
                        </div>
                    </div>

                    <div class="item-square-card">
                        <div class="item-header">
                            <span class="item-tag">7 nap</span>
                        </div>
                        <div class="item-image">
                            <img
                                src="https://dunb17ur4ymx4.cloudfront.net/packages/images/722e66d738b39cb06138462a8e8bfda62f299d2b.png"
                                alt="Arany"
                            />
                        </div>
                        <div class="item-name">
                            <h4>Arany</h4>
                            <span class="item-price">5.69 EUR</span>
                        </div>
                    </div>

                    <div class="item-square-card">
                        <div class="item-header">
                            <span class="item-tag">7 nap</span>
                        </div>
                        <div class="item-image">
                            <img
                                src="https://dunb17ur4ymx4.cloudfront.net/packages/images/5c268d8b4d7984c2e727e564e8f489ce14b9a35e.png"
                                alt="Ezüst"
                            />
                        </div>
                        <div class="item-name">
                            <h4>Ezüst</h4>
                            <span class="item-price">3.69 EUR</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <aside class="sidebar">
            <div class="status-widget">
                <div class="status-container">
                    <div class="status-indicator"></div>
                    <div class="status-info">
                        <span class="status-text">Online</span>
                        <span class="player-count"
                            >{currentPlayers} / {maxPlayers}</span
                        >
                    </div>
                </div>

                <div class="progress-bar">
                    <div
                        class="progress-bar-fill"
                        style="width: {(currentPlayers / maxPlayers) * 100}%"
                    ></div>
                </div>

                <div class="stats">
                    <div class="stat-box">
                        <span class="stat-title">Adminok</span>
                        <span class="stat-value">12</span>
                    </div>
                    <div class="stat-box">
                        <span class="stat-title">Várólistán</span>
                        <span class="stat-value">1000</span>
                    </div>
                </div>
            </div>
        </aside>
    </div>
</main>

<style>
    main {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: var(--width, 240px) 1fr;
        grid-template-rows: var(--height, 48px) 1fr;
        overflow: hidden;
        color: var(--text);
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
    }

    #home-view {
        grid-column: 2;
        grid-row: 2;
        padding: 32px;
        display: grid;
        grid-template-columns: 1fr 280px;
        gap: 32px;
        overflow-y: auto;
    }

    .feed-layout {
        display: flex;
        flex-direction: column;
        gap: 32px;
    }

    .hero-card {
        position: relative;
        min-height: 200px;
        border-radius: 24px;
        background: url("https://news.see-rpg.com/img/YhWW8icjGa7KbAtF3j1V.png")
            center/cover;
        border: 1px solid var(--border);
        overflow: hidden;
        display: flex;
        align-items: flex-end;
        padding: 48px;
        box-shadow: 0 20px 40px -15px rgba(0, 0, 0, 0.6);
    }

    .hero-overlay {
        position: absolute;
        inset: 0;
        background: linear-gradient(
            to top,
            rgba(11, 15, 23, 0.8) 40%,
            rgba(11, 15, 23, 0.4) 100%
        );
        z-index: 1;
    }

    .hero-content {
        position: relative;
        z-index: 2;
        max-width: 600px;
    }

    .badge {
        background: var(--green);
        color: var(--black);
        font-size: 12px;
        padding: 4px 8px;
        border-radius: 8px;
        text-transform: uppercase;
        letter-spacing: 0.5px;
        display: inline-block;
        margin-bottom: 12px;
        font-weight: bold;
    }

    .hero-content h2 {
        margin: 0 0 8px 0;
        font-size: 24px;
    }

    .hero-content p {
        margin: 0 0 20px 0;
        color: var(--comment);
    }

    .hero-news-button {
        background: var(--orange);
        color: var(--black);
        border: none;
        padding: 10px 20px;
        border-radius: 12px;
        cursor: pointer;
        font-size: 14px;
    }

    .status-widget {
        background: var(--surface1);
        border: 1px solid var(--border);
        border-radius: 16px;
        padding: 24px;
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .status-container {
        display: flex;
        align-items: center;
        gap: 12px;
    }

    .status-indicator {
        width: 10px;
        height: 10px;
        background: var(--green);
        border-radius: 50%;
    }

    .status-info {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-grow: 1;
    }

    .status-text {
        font-size: 16px;
        color: var(--green);
    }

    .player-count {
        font-size: 14px;
        color: var(--text);
    }

    .progress-bar {
        width: 100%;
        height: 6px;
        background: rgba(255, 255, 255, 0.05);
        border-radius: 10px;
        overflow: hidden;
    }

    .progress-bar-fill {
        height: 100%;
        background: var(--green);
        border-radius: 10px;
    }

    .stats {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 12px;
        margin-top: 4px;
    }

    .stat-box {
        background: rgba(255, 255, 255, 0.02);
        border: 1px solid var(--border);
        padding: 12px;
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .stat-title {
        font-size: 12px;
        color: var(--text);
        text-transform: uppercase;
    }

    .stat-value {
        font-size: 14px;
        color: var(--comment);
    }

    .items {
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .items-title {
        margin: 0;
        font-size: 18px;
        text-transform: uppercase;
        letter-spacing: 1.5px;
        color: var(--text);
    }

    .items-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
        gap: 16px;
    }

    .item-square-card {
        background: var(--surface1);
        border: 1px solid var(--border);
        border-radius: 14px;
        padding: 16px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        aspect-ratio: 1 / 1;
        box-sizing: border-box;
        cursor: pointer;
    }

    .item-tag {
        font-size: 10px;
        color: var(--text);
        background: rgba(255, 255, 255, 0.06);
        padding: 4px 8px;
        border-radius: 8px;
    }

    .item-image {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-grow: 1;
        margin: 8px 0;
    }

    .item-image img {
        max-height: 65px;
        max-width: 100%;
        object-fit: contain;
        filter: drop-shadow(0 2px 2px rgba(0, 0, 0, 0.5));
        transition: transform 0.2s;
    }

    .item-square-card:hover .item-image img {
        transform: scale(1.2);
    }

    .item-name {
        display: flex;
        flex-direction: column;
        gap: 2px;
    }

    .item-name h4 {
        margin: 0;
        font-size: 14px;
        color: var(--text);
    }

    .item-price {
        font-size: 12px;
        color: var(--green);
    }

    @media (max-width: 900px) {
        #home-view {
            grid-template-columns: 1fr;
        }
        .sidebar {
            order: 0;
        }
    }
</style>
