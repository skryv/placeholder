# placeholder

A placeholder application used for scaffolding ECS Fargate services that does just enough to pass health checks.  

## usage

```
FROM public.ecr.aws/skryv/placeholder:latest
```

You can set the port that the application runs on by setting the `PORT` environment variable.

## references

- https://docs.aws.amazon.com/AmazonECR/latest/public/getting-started-cli.html
