# Todoist CSV Import/Export Guidelines

---

## Overview

Todoist lets you import or export project templates using a CSV file.  
Useful for cloning projects, migrating from other tools, or building reusable task templates.

---

## CSV Structure

Each row in the CSV represents a task, section, or note.  
All headers below must exist and be lowercase.

| Column              | Example                       | Description                                                                  |
| :------------------ | :---------------------------- | :--------------------------------------------------------------------------- |
| **TYPE**            | `task` \| `section` \| `note` | Defines row type. Must be lowercase.                                         |
| **CONTENT**         | `Fix backend bug @urgent`     | Task name, section title, or note content. Supports labels via `@labelname`. |
| **DESCRIPTION**     | `Reproduce issue in staging`  | Task description (optional).                                                 |
| **PRIORITY**        | `1`–`4`                       | `1 = p1 (highest)`, `4 = p4 (lowest)`. Empty → p1.                           |
| **INDENT**          | `1`–`4`                       | 1 = parent, 2 = sub-task, 3 = sub-sub-task. Empty → no indent.               |
| **AUTHOR**          | `Evan (14781400)`             | Creator of the task. Empty → you.                                            |
| **RESPONSIBLE**     | `Rikke (20981802)`            | Assigned user. Empty → none.                                                 |
| **DATE**            | `every monday 10 am`          | Task due date or recurrence rule.                                            |
| **DATE_LANG**       | `en`                          | Language code for date parsing.                                              |
| **TIMEZONE**        | `US/Eastern`                  | Sets time zone. Empty → auto-detect.                                         |
| **DURATION**        | `90`                          | Estimated task duration in minutes.                                          |
| **DURATION_UNIT**   | `minute` \| `None`            | Unit for duration field.                                                     |
| **meta view_style** | `board`                       | Imports as a board layout.                                                   |
| **DEADLINE**        | `2025-12-31`                  | Optional deadline date.                                                      |
| **DEADLINE_LANG**   | `en`                          | Language code for deadline parsing.                                          |

Supported language codes:  
`cs, da, de, en, es, fi, fr, it, ja, ko, nb, nl, pl, pt_BR, ru, sv, zh_CN, zh_TW`

---

## Assigning Tasks

To assign tasks automatically on import:

1. Open the shared Todoist project.
2. Click ⋮ → Export as CSV.
3. Open the file.
4. Copy the assignee’s username (ID) from the RESPONSIBLE cell.
5. Paste it into your new CSV’s RESPONSIBLE column.

---

## Importing a CSV

1. Log in to https://todoist.com.
2. Create or open a project.
3. Click ⋮ → Import from CSV.
4. Upload the file (drag-and-drop or Upload from computer).

Todoist automatically creates tasks, sections, and comments.  
Your project layout (List, Board, Week view) is preserved.

---

## Exporting a CSV

1. Log in to https://todoist.com.
2. Open an existing project.
3. Click ⋮ → Export as CSV.
4. Todoist downloads the project clone, including task data, dates, and attachments.

---

## Troubleshooting

### Encoding Must Be UTF-8

Todoist rejects files not encoded in UTF-8.

**Excel**

```
File → Save As → CSV UTF-8 (Comma delimited)
```

**Numbers (Mac)**

```
File → Export To → CSV → Advanced Options → Text Encoding → Unicode (UTF-8)
```

**Notepad**

```
File → Save As → Encoding: UTF-8 → Name: something.csv
```

---

### Garbled Columns (Excel Fix)

If everything lands in one column:

1. Select column A
2. Data → Text to Columns
3. Choose Delimited → Next
4. Check Semicolon only
5. Finish

---

### Missing Tasks After Import

Todoist skips any row with invalid or mismatched data.

Check:

- TYPE is lowercase (task, section, note)
- All required columns exist
- CSV encoding is UTF-8

---

### Undoing Imports

There’s no built-in undo.  
Select → Delete the imported tasks manually.

---

## Limitations

- Completed tasks aren’t included in exports.  
  Use Export to Google Sheets if you need them (labels/comments excluded).
- Recurring start dates are lost — only the recurrence rule remains.

---

## Localization

Project templates support:

English, Spanish, German, Dutch, Italian, Polish, Turkish, Japanese, Portuguese, Czech, French, Korean, Russian

Templates in any other language default to English.

---

## Support

If import/export formatting still breaks, contact Todoist Support —  
Rikke, Marija, Dermot, or another team member will look at your file.
