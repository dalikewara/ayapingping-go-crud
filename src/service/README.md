- Service is your main business logic
- You should always call your repository methods in this package
- You may use your `src/library` functions directly in this package
- Any changes outside this package should not affect your services (except changes for business entity or repository)
- If you need config variables, external clients, or repositories, pass/inject them as dependency
- You can use your own style as long as it doesn't break the main idea