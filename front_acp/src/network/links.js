export function addLink(graph, id, input, output) {
  var link = new joint.shapes.devs.Link({
    attrs: {
      acp_id: id,
    },
    source: {
      id: output.device_id,
      port: output.port
    },
    target: {
      id: input.device_id,
      port: input.port
    }
  });

  link.addTo(graph)
  return link
}
