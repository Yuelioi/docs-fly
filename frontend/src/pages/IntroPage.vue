<template>
    <Panel
        header="Header"
        toggleable
        :p="{
            header: (options) => ({
                id: 'myPanelHeader',
                style: {
                    'user-select': 'none'
                },
                class: [
                    'border-primary',
                    {
                        'bg-primary text-primary-contrast': options.state.d_collapsed,
                        'text-primary bg-primary-contrast': !options.state.d_collapsed
                    }
                ]
            }),
            content: { class: 'border-primary text-lg text-primary-700' },
            title: 'text-xl', // OR { class: 'text-xl' }
            toggler: () =>
                'bg-primary text-primary-contrast hover:text-primary hover:bg-primary-contrast' // OR { class: 'bg-primary text-primary-contrast hover:text-primary hover:bg-primary-contrast' }
        }">
        <p class="m-0">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
            incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
            exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure
            dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
            Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
            mollit anim id est laborum.
        </p>
    </Panel>
</template>

<script setup></script>
