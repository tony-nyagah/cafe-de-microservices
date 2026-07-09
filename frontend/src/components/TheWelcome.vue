<script setup lang="ts">
import WelcomeItem from "./WelcomeItem.vue";
import DocumentationIcon from "./icons/IconDocumentation.vue";
import ToolingIcon from "./icons/IconTooling.vue";
import EcosystemIcon from "./icons/IconEcosystem.vue";
import CommunityIcon from "./icons/IconCommunity.vue";
import SupportIcon from "./icons/IconSupport.vue";

const openReadmeInEditor = () => fetch("/__open-in-editor?file=README.md");
</script>

<template>
  <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
    <WelcomeItem>
      <template #icon>
        <DocumentationIcon />
      </template>
      <template #heading>Documentation</template>

      Vue's
      <a href="https://vuejs.org/" target="_blank" rel="noopener" class="link link-primary"
        >official documentation</a
      >
      provides you with all information you need to get started.
    </WelcomeItem>

    <WelcomeItem>
      <template #icon>
        <ToolingIcon />
      </template>
      <template #heading>Tooling</template>

      This project is served and bundled with
      <a
        href="https://vite.dev/guide/features.html"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >Vite</a
      >. The recommended IDE setup is
      <a
        href="https://code.visualstudio.com/"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >VSCode</a
      >
      +
      <a
        href="https://github.com/vuejs/language-tools"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >Vue - Official</a
      >. If you need to test your components and web pages, check out
      <a href="https://vitest.dev/" target="_blank" rel="noopener" class="link link-primary"
        >Vitest</a
      >
      and
      <a href="https://www.cypress.io/" target="_blank" rel="noopener" class="link link-primary"
        >Cypress</a
      >
      /
      <a href="https://playwright.dev/" target="_blank" rel="noopener" class="link link-primary"
        >Playwright</a
      >.

      <br />

      More instructions are available in
      <a href="javascript:void(0)" @click="openReadmeInEditor"><code>README.md</code></a
      >.
    </WelcomeItem>

    <WelcomeItem>
      <template #icon>
        <EcosystemIcon />
      </template>
      <template #heading>Ecosystem</template>

      Get official tools and libraries for your project:
      <a href="https://pinia.vuejs.org/" target="_blank" rel="noopener" class="link link-primary"
        >Pinia</a
      >,
      <a href="https://router.vuejs.org/" target="_blank" rel="noopener" class="link link-primary"
        >Vue Router</a
      >,
      <a
        href="https://test-utils.vuejs.org/"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >Vue Test Utils</a
      >, and
      <a
        href="https://github.com/vuejs/devtools"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >Vue Dev Tools</a
      >. If you need more resources, we suggest paying
      <a
        href="https://github.com/vuejs/awesome-vue"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >Awesome Vue</a
      >
      a visit.
    </WelcomeItem>

    <WelcomeItem>
      <template #icon>
        <CommunityIcon />
      </template>
      <template #heading>Community</template>

      Got stuck? Ask your question on
      <a href="https://chat.vuejs.org" target="_blank" rel="noopener" class="link link-primary"
        >Vue Land</a
      >
      (our official Discord server), or
      <a
        href="https://stackoverflow.com/questions/tagged/vue.js"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >StackOverflow</a
      >. You should also follow the official
      <a
        href="https://bsky.app/profile/vuejs.org"
        target="_blank"
        rel="noopener"
        class="link link-primary"
        >@vuejs.org</a
      >
      Bluesky account or the
      <a href="https://x.com/vuejs" target="_blank" rel="noopener" class="link link-primary"
        >@vuejs</a
      >
      X account for latest news in the Vue world.
    </WelcomeItem>

    <WelcomeItem>
      <template #icon>
        <SupportIcon />
      </template>
      <template #heading>Support Vue</template>

      As an independent project, Vue relies on community backing for its sustainability. You can
      help us by
      <a href="https://vuejs.org/sponsor/" target="_blank" rel="noopener" class="link link-primary"
        >becoming a sponsor</a
      >.
    </WelcomeItem>
  </div>
</template>
