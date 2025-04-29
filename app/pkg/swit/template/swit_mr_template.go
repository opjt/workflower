package template

const MRTemplate = `
{
  "header": {
    "title": "Merge Request: {{.Title}}",
    "subtitle": "Author: {{.AuthorName}} ({{.AuthorUsername}})",
    "app_id": "{{.AppID}}",
    "icon": {
      "type": "image",
      "image_url": "https://images.ctfassets.net/xz1dnu24egyd/1IRkfXmxo8VP2RAE5jiS1Q/ea2086675d87911b0ce2d34c354b3711/gitlab-logo-500.png",
      "alt": "GitLab Logo"
    }
  },
  "body": {
    "elements": [
      {
        "type": "collection_entry",
        "text_sections": [
          {
            "text": {
              "type": "text",
              "content": "{{.MRTitle}}",
              "style": {
                "bold": true,
                "color": "gray800",
                "size": "medium"
              }
            },
            "metadata_items": []
          }
        ],
        "start_section": {
          "type": "image",
          "image_url": "{{.AuthorAvatarURL}}",
          "alt": "{{.AuthorName}} Avatar",
          "style": {
            "size": "small"
          }
        },
        "vertical_alignment": "middle",
        "background": {
          "color": "lightblue"
        },
        "action_id": "action_collection_entry",
        "static_action": {
          "action_type": "open_link",
          "link_url": "{{.MRURL}}"
        },
        "draggable": false
      },
      {
        "type": "collection_entry",
        "text_sections": [
          {
            "text": {
              "type": "text",
              "content": "**Project**: {{.ProjectName}} \n **Branch**: {{.SourceBranch}} ➡️ {{.TargetBranch}}",
              "markdown": true,
              "style": {
                "color": "gray800",
                "size": "medium"
              }
            },
            "metadata_items": [
              {
                "type": "tag",
                "content": "{{.MRState}}",
                "style": {
                  "color": "primary",
                  "shape": "rounded"
                }
              },
              {
                "type": "subtext",
                "content": "{{.CreatedAt}}"
              }
            ]
          }
        ],
        "vertical_alignment": "middle",
        "draggable": false
      }
    ]
  }
}
`
