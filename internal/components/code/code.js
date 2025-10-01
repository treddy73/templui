(function () {
  "use strict";

  // Prevent multiple initializations
  if (window.__codeInitialized) return;
  window.__codeInitialized = true;

  // Highlight code blocks
  function highlightCode(block) {
    if (block && window.hljs && !block.classList.contains("hljs")) {
      window.hljs.highlightElement(block);
    }
  }

  // Update theme
  function updateTheme() {
    const isDark = document.documentElement.classList.contains("dark");
    const links = document.querySelectorAll("#highlight-theme");
    links.forEach((link) => {
      link.href = isDark
        ? "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/atom-one-dark.min.css"
        : "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/atom-one-light.min.css";
    });
  }

  // Initialize
  updateTheme();

  // Simple observer just for dark class
  const observer = new MutationObserver(() => {
    updateTheme();
    // Also highlight any new blocks
    document
      .querySelectorAll("[data-tui-code-block]:not(.hljs)")
      .forEach(highlightCode);
  });

  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ["class"],
  });

  // Also observe body for new content
  const bodyObserver = new MutationObserver(() => {
    document
      .querySelectorAll("[data-tui-code-block]:not(.hljs)")
      .forEach(highlightCode);
  });

  bodyObserver.observe(document.body, {
    childList: true,
    subtree: true,
  });

  // Initial highlight
  document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll("[data-tui-code-block]").forEach(highlightCode);
  });
})();
