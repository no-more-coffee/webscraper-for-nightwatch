# Nightwatch web scraper

## Task

1. First:
    - Go to http://google.com.
    - Search for "pizza".
    - ~~Set the preferences to get 100 results~~ (feature removed).
    - Save the HTML.
2. Create a Go program that does the following:
    - Based on CSS selector definitions read from a local JSON file, extract the following information:
        - URL rank type (organic, local, carousel, knowledge panel, featured snippet),
        - Website title, URL, description, and rank position within its own rank type group.
    - Format the results in a standard format (JSON or YAML, **or CSV**).
    - The use of well-maintained and mature external libraries is encouraged.
    - Ensure optimal memory and CPU usage.
    - Results should be consistent, even when one or more result types are missing.
    - Bonus points for high performance.

## Usage

```bash
go run .
```

- Input HTML is read from [pizza.html](pizza.html).
- Selector groups are read from the [group-selectors.json](group-selectors.json).
- Output is saved to [result.csv](out/result.csv).

### Selector groups

Expected JSON format:

```json
{
   "<group name>": {
      "base": "<selector for group item>",
      "title": "<from item, selected by 'base' selector",
      "url": "<from item, selected by 'base' selector",
      "description": "<from item, selected by 'base' selector"
   }
}
```

Example:

```html

<items-group1>
   <item>
      <title></title>
      <url></url>
      <description></description>
   </item>
   <item>
      ...
   </item>
</items-group1>
```
```json
{
   "organic": {
      "base": "items-group1 > item",
      "title": "title",
      "url": "url",
      "description": "description"
   }
}
```
