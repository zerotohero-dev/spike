```txt
  \\ 
 \\\\ SPIKE: Keep your secrets secret with SPIFFE.
\\\\\\
```

## Quickstart

Make sure you have [SPIRE](https://spiffe.io/spire) installed on your system.

The `./hack/build-spire.sh` can be used as a starting point to build SPIRE
from source and install it to your system.

Set up environment variables

```bash
# This is the SPIKE repo folder that you cloned.
export SPIKE_ROOT=/path/to/the/project/folder/of/spike/
```

Start SPIRE Server:

```bash
./hack/start-spire-server.sh
```

Create a join token:

```bash 
./hack/generate-agent-token.sh
```

Register SPIKE apps:

```bash
./hack/register-spire-entries.sh
```

(Optional) verify registration:

```bash
./hack/show-spire-server-entries.sh
```

Start SPIRE Agent:

```bash 
./hack/start-spire-agent.sh
```

Start the workloads:

```bash
./nexus
./keeper
./pilot
```

That's about it.

Enjoy.


```text
TODO: convert this to an ADR

Idea: 
Issue management:
* This is a tiny project; so it does not need a big fat issue manager.
  even a `to_do.txt` with every line in priority order is a good enough way
  to manage things.
* The development team (me, Volkan, initially) will use `to do` labels liberally
  to designate what to do where in the project.
* GitHub issues will be created on a "per need" basis. 
* Also the community will be encouraged to create GitHub issues, yet it won't
  be the team's main way to define issues or roadmap.
* I believe this unorthodox way will provide agility.


```