# 🎧 kubectl-lofi

**Play lo-fi beats in your terminal while you kubectl.**

> Because nothing says "I'm in control of this 47-node cluster" like some chill hip-hop beats.

---

kubectl-lofi is a [krew](https://krew.sigs.k8s.io/) plugin that streams lo-fi radio directly into your terminal. That's it. That's the whole thing.

Born from the scientifically unproven but spiritually undeniable fact that **lo-fi music makes you mass better at Kubernetes** *(citation needed)*. Debugging CrashLoopBackOff at 2 AM? Lo-fi. Rolling back a bad deployment in prod? Lo-fi. Watching your nodes get evicted one by one? *Definitely* lo-fi.

## Installation

### 1. Install krew (if you haven't already)

```bash
brew install krew
```

### 2. Install the plugin

```bash
kubectl krew install --manifest-url=https://raw.githubusercontent.com/zamai/kubectl-lofi/main/lofi.yaml
```

### 3. Vibe

```bash
kubectl lofi
```

```
♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫
♫      _          _        _   _       ♪
♪     | | ___   _| |__  __| |_| |      ♫
♫     | |/ / | | | '_ \/ _|  _| |      ♪
♪     |   <| |_| | |_) \__|_| |_|      ♫
♫     |_|\_\\__,_|_.__/               ♪
♪           _          __  _            ♫
♫          | |  ___   / _|(_)           ♪
♪          | | / _ \ |  _|| |           ♫
♫          |_| \___/ |_|  |_|           ♪
♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫  ♪  ♫

🎧 Now playing lo-fi radio...
Press Ctrl+C to stop.
```

Press `Ctrl+C` when you've achieved inner peace (or the incident is resolved, whichever comes first).

## Why?

Because Kubernetes is stressful and you deserve nice things.

### Scientifically Questionable Benefits

| K8s Operation | Without Lo-fi | With Lo-fi |
|---|---|---|
| `kubectl delete pod` | Anxiety | Zen |
| `helm upgrade --install` | Sweaty palms | Smooth vibes |
| Debugging OOMKilled | Existential dread | Mild inconvenience |
| `kubectl drain node` | Panic | Groovy |
| Reading YAML | Pain | Slightly less pain |
| 3 AM PagerDuty alert | Rage quit | *sips coffee aesthetically* |

## Requirements

- macOS (darwin/amd64 or darwin/arm64)
- Audio output (headphones recommended for maximum aesthetic)
- A Kubernetes cluster to pretend to manage while vibing

## How It Works

It streams MP3 audio. In your terminal. While you kubectl. There is no AI, no blockchain, no machine learning. Just vibes.

## FAQ

**Q: Does this actually improve my Kubernetes skills?**
A: Legally, we cannot confirm or deny this. But also yes.

**Q: Can I use this in production?**
A: This *is* production. The lo-fi never stops in production.

**Q: Why macOS only?**
A: Because Mac users need the most emotional support while doing DevOps.

**Q: My cluster is on fire, should I start lo-fi first or fix it first?**
A: Lo-fi first. Always lo-fi first. You can't fix anything in a state of panic.

**Q: Does it work with OpenShift?**
A: It works with anything that makes you stressed, so yes, especially OpenShift.

## Contributing

If you want to add Linux support, Windows support, or more streams — PRs are welcome. If you want to add a feature that isn't vibes-related, please reconsider.

## License

[MIT](LICENSE) — Free as in "free to vibe."

---

<p align="center">
  <i>Made with 🎵 for the mass who keeps Kubernetes running at unreasonable hours.</i>
</p>
