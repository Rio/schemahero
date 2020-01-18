package installer

var generatedMigrationCRDV1Beta1 = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.4
  creationTimestamp: null
  name: migrations.schemas.schemahero.io
spec:
  group: schemas.schemahero.io
  names:
    kind: Migration
    listKind: MigrationList
    plural: migrations
    singular: migration
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: Migration is the Schema for the migrations API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: MigrationSpec defines the desired state of Migration
          type: object
        status:
          description: MigrationStatus defines the observed state of Migration
          type: object
      type: object
  version: v1alpha3
  versions:
  - name: v1alpha3
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
var generatedMigrationCRDV1 = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.4
  creationTimestamp: null
  name: migrations.schemas.schemahero.io
spec:
  group: schemas.schemahero.io
  names:
    kind: Migration
    listKind: MigrationList
    plural: migrations
    singular: migration
  scope: Namespaced
  versions:
  - name: v1alpha3
    schema:
      openAPIV3Schema:
        description: Migration is the Schema for the migrations API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: MigrationSpec defines the desired state of Migration
            type: object
          status:
            description: MigrationStatus defines the observed state of Migration
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
