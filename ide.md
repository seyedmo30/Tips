چیزایی که باید رو vscode نصب شه :

Git Graph  ----   ext install mhutchie.git-graph

Error Lens  ---   ext install usernamehw.errorlens

Paste JSON as Code --- ext install quicktype.quicktype

Prettier - Code formatter --- ext install esbenp.prettier-vscode

bookmarks - برای علامت زدن روی کدی که مهمه در یه فایل بزرگ

## keymap shortcuts

+ go to definition (ctrl + click)

میره جایی که تابع تعریف شده

+ go to impelemention  (ctrl + shift + space)

اگر اینترفیس باشه جایی که اون رو کانکریت کرده

+ go to refrence  (ctrl + click)

میره جاهایی که از این فانکشن استفاده شده

+ go to symbol

نمایش فانکشن ها ، استراکت ها و کلاس ها در صفحه

+ peek definition (ctrl + space)

برای دسترسی به پارامتر های تابع

+ go back (alt + leftArrow)

کار قبلی که انجام دادیم

+ go forward (alt + rightArrow)

بازگشت به کاری که انجام دادیم




#### home/seyed/.config/Code/User/keybindings.json

```


// Place your key bindings in this file to override the defaultsauto[]
[
    {
        "key": "alt+left",
        "command": "-workbench.action.terminal.focusPreviousPane",
        "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported"
    },
    {
        "key": "alt+left",
        "command": "-gitlens.key.alt+left",
        "when": "gitlens:key:alt+left"
    },
    {
        "key": "alt+left",
        "command": "workbench.action.navigateBack",
        "when": "canNavigateBack"
    },
    {
        "key": "ctrl+alt+-",
        "command": "-workbench.action.navigateBack",
        "when": "canNavigateBack"
    },
    {
        "key": "alt+right",
        "command": "-workbench.action.terminal.focusNextPane",
        "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported"
    },
    {
        "key": "alt+right",
        "command": "-gitlens.key.alt+right",
        "when": "gitlens:key:alt+right"
    },
    {
        "key": "alt+right",
        "command": "workbench.action.navigateForward",
        "when": "canNavigateForward"
    },
    {
        "key": "ctrl+shift+-",
        "command": "-workbench.action.navigateForward",
        "when": "canNavigateForward"
    },
    {
        "key": "ctrl+shift+space",
        "command": "editor.action.peekDefinition",
        "when": "editorHasDefinitionProvider && editorTextFocus && !inReferenceSearchEditor && !isInEmbeddedEditor"
    },
    {
        "key": "ctrl+shift+f10",
        "command": "-editor.action.peekDefinition",
        "when": "editorHasDefinitionProvider && editorTextFocus && !inReferenceSearchEditor && !isInEmbeddedEditor"
    }
]


```