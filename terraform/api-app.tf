variable "environment" {}
variable "api_image" {}

resource "kubernetes_manifest" "cloudrun-api" {
  manifest = {
    "apiVersion" : "serving.knative.dev/v1",
    "kind" : "Service",
    "metadata" : {
      "name" : "api-app-${var.environment}",
      // "namespace" : "default",
      "namespace" : var.environment,
      "labels" : {
        "networking.knative.dev/visibility" : "cluster-local"
      },
    },
    "spec" : {
      "template" : {
        "spec" : {
          "serviceAccountName" : "default",
          "timeoutSeconds" : 300,
          "containers" : [
            {
              "name" : "user-container",
              "image" : var.api_image,
              "imagePullPolicy": "Always",
              "ports" : [
                {
                  "containerPort" : 8080,
                  "protocol" : "TCP"
                }
              ],
              "env" : [
                {
                  "name" : "API_PORT",
                  "value" : "8080"
                },
                {
                  "name" : "PROJECT_ID",
                  "value" : var.project
                }
              ],
              "resources" : {
                "limits" : {
                  "memory" : "256Mi"
                }
              },
              "readinessProbe" : {
                "successThreshold" : 1,
                "tcpSocket" : {}
              }
            }
          ]
        }
      },
      "traffic" : [
        {
          "percent" : 100,
          "latestRevision" : true
        }
      ],
    }
  }
}
