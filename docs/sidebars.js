/**
 * Copyright (c) 2017-present, Facebook, Inc.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

module.exports = {
  adminSidebar: [
    {
      type: "doc",
      id: "what-is-devpod",
    },
    {
      type: "category",
      label: "Getting Started",
      collapsed: false,
      items: [
        {
          type: "doc",
          id: "getting-started/quickstart",
        },
      ],
    },
    {
      type: "category",
      label: "Developing in a Workspace",
      items: [
        {
          type: "doc",
          id: "developing-in-workspaces/what-are-workspaces",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/create-a-workspace",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/devcontainer-json",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/prebuild-a-workspace",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/credentials",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/stop-a-workspace",
        },
        {
          type: "doc",
          id: "developing-in-workspaces/delete-a-workspace",
        },
      ],
    },
    {
      type: "category",
      label: "Managing your Machines",
      items: [
        {
          type: "doc",
          id: "managing-machines/what-are-machines",
        },
        {
          type: "doc",
          id: "managing-machines/manage-machines",
        },
      ],
    },
    {
      type: "category",
      label: "Managing your Providers",
      items: [
        {
          type: "doc",
          id: "managing-providers/what-are-providers",
        },
        {
          type: "doc",
          id: "managing-providers/add-provider",
        },
        {
          type: "doc",
          id: "managing-providers/update-provider",
        },
        {
          type: "doc",
          id: "managing-providers/delete-provider",
        },
        {
          type: "doc",
          id: "managing-providers/docker",
        },
        {
          type: "doc",
          id: "managing-providers/kubernetes",
        },
      ],
    },
    {
      type: "category",
      label: "Developing Providers",
      items: [
        {
          type: "doc",
          id: "developing-providers/quickstart",
        },
        {
          type: "doc",
          id: "developing-providers/options",
        },
        {
          type: "doc",
          id: "developing-providers/binaries",
        },
        {
          type: "doc",
          id: "developing-providers/agent",
        },
      ],
    },
    {
      type: "link",
      label: "Originally created by Loft",
      href: "https://loft.sh/",
    },
  ],
};
