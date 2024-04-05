<div tabs name="Create a Subscription" group="create-subscription">
  <details open>
  <summary label="Kyma Dashboard">
  Kyma Dashboard
  </summary>

1. Go to **Namespaces** and select the default namespace.
2. Go to **Configuration** > **Subscriptions** and click **Create Subscription+**.
3. Switch to the **Advanced** tab, and provide the following parameters:
   - **Subscription name**: `lastorder-sub`
   - **Config**: `maxInFlightMessages: 5`
   - **Types**: `order.received.v1`
   - **Service**: `lastorder` (The sink field will be populated automatically.)
   - **Type matching:**: `standard`
   - **Source**: `myapp`

5. Click **Create**.
6. Wait a few seconds for the Subscription to have status `READY`.

</details>
<details>
<summary label="kubectl">
kubectl
</summary>

Run:

cat <<EOF | kubectl apply -f -
    apiVersion: eventing.kyma-project.io/v1alpha2
    kind: Subscription
    metadata:
      name: lastorder-sub
      namespace: default
    spec:
      sink: 'http://lastorder.default.svc.cluster.local'
      source: myapp
      types:
       - order.received.v1
       - order.changed.v1
EOF

To check that the Subscription was created and is ready, run: `kubectl get subscriptions lastorder-sub -o=jsonpath="{.status.ready}"

The operation was successful if the returned status says `true`.

  </details>
</div>






2. Deploy an instance of the HttpBin service or a sample Function.

  <div tabs name="create">

    <details>
    <summary>
    HttpBin
    </summary>

    To deploy an instance of the HttpBin service in your namespace using the [sample code](https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/httpbin.yaml), run:

    ```shell
    kubectl -n $NAMESPACE create -f https://raw.githubusercontent.com/istio/istio/master/samples/httpbin/httpbin.yaml
    ```

    </details>

    <details>
    <summary>
    Function
    </summary>

    To create a Function in your namespace using the [sample code](https://raw.githubusercontent.com/kyma-project/kyma/main/docs/03-tutorials/00-api-exposure/assets/function.yaml), run:

    ```shell
    kubectl -n $NAMESPACE apply -f https://raw.githubusercontent.com/kyma-project/kyma/main/docs/03-tutorials/00-api-exposure/assets/function.yaml
    ```

    </details>
  </div>

3. Verify if an instance of the HttpBin service or a sample Function is successfully created.

  <div tabs name="verify">

    <details>
    <summary>
    HttpBin
    </summary>

    To verify if an instance of the HttpBin service is created, run:

      ```shell
        kubectl get pods -l app=httpbin -n $NAMESPACE
      ```

    You should get a result similar to this one:

      ```shell
        NAME             READY    STATUS     RESTARTS    AGE
        httpbin-test     2/2      Running    0           96s
      ```

    </details>

    <details>
    <summary>
    Function
    </summary>

    To verify if a Function is created, run:

      ```shell
        kubectl get functions $NAME -n $NAMESPACE
      ```

    You should get a result similar to this one:

      ```shell
        NAME            CONFIGURED   BUILT     RUNNING   RUNTIME    VERSION   AGE
        test-function   True         True      True      nodejs18   1         96s
      ```
    </details>
  </div>


# Test Infra
<!-- markdown-link-check-disable-next-line -->
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkyma-project%2Ftest-infra.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkyma-project%2Ftest-infra?ref=badge_shield)
<!-- markdown-link-check-disable-next-line -->
[![REUSE status](https://api.reuse.software/badge/github.com/kyma-project/test-infra)](https://api.reuse.software/info/github.com/kyma-project/test-infra)

## Overview

The purpose of the `test-infra` repository is to store configuration and scripts for the test infrastructure used in the `kyma-project` organization.
See also the internally available [test-infra onboarding](https://github.tools.sap/kyma/test-infra/blob/main/onboarding.md).

### Project Documentation

Please see the [index page](/docs/index.md) for the Test Infra documentation. It lists all the documentation available in the `test-infra` repository.

### Prow

The `test-infra` repository contains the whole configuration of Prow. Its purpose is to replace the internal Continuous Integration (CI) tool in the `kyma-project` organization.

For more detailed documentation on installation, configuration, development, and testing, go to the [`docs/prow`](./docs/prow) subfolder.


## License
<!-- markdown-link-check-disable-next-line -->
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkyma-project%2Ftest-infra.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkyma-project%2Ftest-infra?ref=badge_large)

## Contributing
<!--- mandatory section - do not change this! --->

See the [Contributing Rules](CONTRIBUTING.md).

## Code of Conduct
<!--- mandatory section - do not change this! --->

See the [Code of Conduct](CODE_OF_CONDUCT.md) document.

## Licensing
<!--- mandatory section - do not change this! --->

See the [license](./LICENSE) file.
