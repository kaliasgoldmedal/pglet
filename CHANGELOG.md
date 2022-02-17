# Change Log - Pglet Server

## [0.7.0](https://github.com/pglet/pglet/releases/tag/v0.7.0)

New `SplitStack` control (based on [split.js](https://split.js.org/)) added with the following properties:
* `horizontal` (bool)
* `gutterSize` (int)
* `gutterColor` (string)
* `gutterHoverColor` (string)
* `gutterDragColor` (string)
* `resize` (event)

New `TextBox` properties:
* `shiftEnter` (bool) - blocks ENTER button in `multiline` TextBox, but pops up the event, so `Stack.submit` could be triggered. New line could still be entered with SHIFT+ENTER. This is to build Discord-like message box.
* `rows` (int) - sets initial size in rows of `multiline` TextBox.
* `resizable` (bool) - controls whether `multiline` TextBox is resizable by the user. Default is `true`. `autoAdjustHeight` is still respected even if `resizable` is `false`.

`Panel` control changes:
* `blocking` (bool) is now `true` by default.

## [0.6.0](https://github.com/pglet/pglet/releases/tag/v0.6.0)

* `Stack` automatically scrolls to bottom if `autoscroll` property set to `true`.
* Set `page.UserAuthProvider` to a used authentication method (`github`, `google` or `azure`).
* `page.win_width` and `page.win_height` properties renamed to `page.winwidth` and `page.winheight`.
* When host is connected to a `page` its contents and properties are cleaned unless `update: true` is passed. No need to call `page.clean()` on the client anymore.
* Focusing input controls - allows setting focus on a control when added to a page or page loaded:
  * `Button.focused`
  * `Checkbox.focused`
  * `ChoiceGroup.focused`
  * `DatePicker.focused`
  * `Dropdown.focused`
  * `SearchBox.focused`
  * `Slider.focused`
  * `SpinButton.focused`
  * `Textbox.focused`
  * `Toggle.focused`
* `focus` and `blur` events for the following input controls:
  * `Button`
  * `ChoiceGroup`
  * `DatePicker`
  * `Dropdown`
  * `SearchBox`
  * `Slider`
  * `SpinButton`
  * `Textbox`
  * `Toggle`
* New `page` properties:
  * `userAuthProvider`
* New `page` events:
  * `resize`
  * `connect` - web client connected
  * `disconnect` - web client disconnected
* New `IFrame` properties:
  * `borderWidth`
  * `borderColor`
  * `borderStyle`
  * `borderRadius`
* New `Stack` properties:
  * `autoscroll`
  * `borderWidth`
  * `borderColor`
  * `borderStyle`
* New `Stack` events:
  * `submit`
* New `Image` properties:
  * `fit` = `none`, `contain`, `cover`, `center`, `centerContain`, `centerCover`
  * `borderWidth`
  * `borderColor`
  * `borderStyle`
  * `borderRadius`
* New `Dropdown.Option` properties:
  * `itemType` (`normal`, `divider`, `header`)
  * `disabled`
* New `Persona` control:
  * `imageUrl`
  * `imageAlt`
  * `initialsColor`
  * `initialsTextColor`
  * `text`
  * `secondaryText`
  * `tertiaryText`
  * `optionalText`
  * `size`
  * `presence`
  * `hideDetails`
* New `ComboBox` control:
  * `label`
  * `value`
  * `placeholder`
  * `error_message`
  * `focused`
  * `multi_select`
  * `allow_free_form`
  * `auto_complete`
  * `options`:
    * `key`
    * `text`
    * `itemType` (`normal`, `divider`, `header`, `select_all`)
    * `disabled`

* Removed `IFrame` properties:
  * `border`
* Removed `Stack` properties:
  * `border`
  * `borderLeft`
  * `borderRight`
  * `borderTop`
  * `borderBottom`
* Removed `Text` properties:
  * `borderLeft`
  * `borderRight`
  * `borderTop`
  * `borderBottom`

Bug fixes:

* Duplicate React rendering when loading a page.
