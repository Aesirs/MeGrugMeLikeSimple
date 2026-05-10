/// <reference lib="dom" />

const DARK_KEY = 'dark';

function isDark(): boolean {
  return localStorage.getItem(DARK_KEY) === 'true';
}

function applyDark(checked: boolean): void {
  document.documentElement.classList.toggle('dark', checked);
  localStorage.setItem(DARK_KEY, checked ? 'true' : 'false');
}

/** Call from onclick on the toggle checkbox */
function toggleDark(checked: boolean): void {
  applyDark(checked);
}

// ─── Init (runs synchronously on load to prevent flash) ──
if (isDark()) {
  document.documentElement.classList.add('dark');
}

// ─── DOM ready: sync checkbox state ────────────────────
document.addEventListener('DOMContentLoaded', () => {
  const toggle = document.getElementById('dark-toggle') as HTMLInputElement | null;
  if (toggle) {
    toggle.checked = isDark();
  }
});

// Expose for inline onclick in templates
(window as any).toggleDark = toggleDark;
