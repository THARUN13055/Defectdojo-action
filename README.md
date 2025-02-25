# DefectDojo Scaning Report Upload action

This is the GitHub Action which make you as easly **Upload** your Scanning Reports.

## Features

- **Sending Report files**
- **Automatically Create Engagements**  
- **Golang Updated package**  
- **Single Engagement** It will Store Multiple Scanner Reports

## Use Cases

Here this Github Action is mainly focus on **Sending the Report file** to DefectDojo:

- **Scanner Reports** – Using the Product Name It will Upload the Scanner Reports
- **Auto Deployment** – Here We need to add only the env
- **Deployment Alerts** – Alert when the Report is Uploaded or Not

---

## Env Needed For the Actions

| With            | Description                                        | Required |
|-----------------|------------------------------------------------|----------|
| `DEFECTDOJO_TOKEN`  | Here we need to add the defectdojo **Token**           | ✅       |
| `DEFECTDOJO_URL`  | Url of the Defectdojo **<http://ip:port>**         | ✅       |
| `FILE_PATH_WITH_FILE_NAME` | **Path** of the file and also with **filename** | ✅       |
| `PRODUCT_NAME` | Mension your Specific Product Name which you have been Created | ✅       |
| `SCAN_TYPE`       | Mension the Report Type like **Trivy Scan**                  | ✅      |

---

## Instructions to Use this Action

### Create the API v2 Token

1. Go to Your DefectDojo Page.
2. Go to this path **<http://ip:port/api/key-v2>**
3. Use Old Token or Generate a new token.

### Created a Products

Here you need to Created the User and Products before you running the actions.

- Go to this url **<http://ip:port/product/add>** and add the **Products**

### Templates Files

Inside `.github/workflows/{your-filename}.yml`, add the following workflow:

```yaml
name: DefectDojo

on: [push]

jobs:
  DefectDojo_Report_Upload:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Repository
        uses: actions/checkout@v4
      
      - name: Upload the Report files.
        uses: tharun13055/Defectdojo-action@v1.0.0
        with:
          DEFECTDOJO_TOKEN: ${{ secrets.DEFECTDOJO_TOKEN }}
          DEFECTDOJO_URL: ${{ secrets.DEFECTDOJO_URL }}
          FILE_PATH_WITH_FILE_NAME: ""
          PRODUCT_NAME: ""
          SCAN_TYPE: ""
```
### Refere the FileType.md

For more details on the supported file types, please refer to the [Filetype.md](./FileType.md) file.

## Reports & Feedback

Please report any issues or feature requests via the **[GitHub repository's issue tracker](https://github.com/THARUN13055/Defectdojo-action/issues)**.
