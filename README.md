# Go Web Application - Jenkins Multibranch Pipeline

This repository contains a simple Go web application and a Declarative `Jenkinsfile` configured for Jenkins Multibranch Pipeline with branch-specific execution stages.

---

## 📁 Repository Structure

```
.
├── .gitignore          # Git ignore rules
├── Jenkinsfile         # Jenkins multibranch pipeline definition
├── README.md           # Instructions & documentation
├── go.mod              # Go module file
├── main.go             # Go HTTP server application
└── main_test.go        # Unit tests
```

---

## 🚀 Branch-Specific Pipeline Behavior

| Pipeline Stage | `main` Branch | `dev` / `feature/*` Branch |
| :--- | :--- | :--- |
| **Checkout & Info** | Runs ✅ | Runs ✅ |
| **Compile & Build** | Runs ✅ | Runs ✅ |
| **Unit Tests** | Runs ✅ | Runs ✅ |
| **Dev Deployment & Verification** | **Skipped** ⏭️ | Runs ✅ (`when { branch 'dev' }`) |
| **Production Deployment & Release** | Runs ✅ (`when { branch 'main' }`) | **Skipped** ⏭️ |

---

## 🛠️ Step-by-Step Git Commands

### 1. Initialize Git and Commit to `main`

```bash
git init
git branch -M main
git add .
git commit -m "feat: initial commit with Go web app and Jenkinsfile"
git remote add origin <YOUR_REMOTE_REPOSITORY_URL>
git push -u origin main
```

### 2. Create the `dev` Branch and Push

```bash
git checkout -b dev
# (Optional) Make a change in dev
git push -u origin dev
```

---

## ⚙️ Jenkins Multibranch Pipeline Setup

1. **Create Job in Jenkins**:
   - In Jenkins Dashboard, click **New Item**.
   - Enter item name (e.g., `go-web-app-pipeline`).
   - Select **Multibranch Pipeline** and click **OK**.

2. **Configure Branch Sources**:
   - Under **Branch Sources**, click **Add source** -> choose **Git** (or GitHub / GitLab).
   - Enter your repository URL and credentials.
   - Under **Discover branches**, select **All branches** (or Filter by name: `main* dev* feature/*`).

3. **Build Configuration**:
   - Mode: `by Jenkinsfile`
   - Script Path: `Jenkinsfile`

4. **Scan & Run**:
   - Click **Save**. Jenkins will automatically trigger **Branch Indexing / Scan Multibranch Pipeline**.
   - It will automatically discover both `main` and `dev` branches, and run their respective pipelines.
