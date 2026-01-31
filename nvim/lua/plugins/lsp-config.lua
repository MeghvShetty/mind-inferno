return {
	-- Mason: installs LSP servers, DAPs, etc.
	{
		"williamboman/mason.nvim",
		config = function()
			require("mason").setup()
		end,
	},

	-- Mason-LSPConfig: bridges Mason and the LSP setup
	{
		"williamboman/mason-lspconfig.nvim",
		config = function()
			require("mason-lspconfig").setup({
				ensure_installed = { "lua_ls", "gopls" },
			})
		end,
	},

	-- Core LSP configuration (new API)
	{
		"neovim/nvim-lspconfig",
		config = function()
			local capabilities = require("cmp_nvim_lsp").default_capabilities()

			-- Define LSP configs
			vim.lsp.config.lua_ls = {
				capabilities = capabilities,
				settings = {
					Lua = {
						diagnostics = { globals = { "vim" } },
					},
				},
			}

			vim.lsp.config.gopls = {
				capabilities = capabilities,
			}

			-- Autostart servers based on filetype
			vim.api.nvim_create_autocmd("FileType", {
				pattern = "lua",
				callback = function()
					vim.lsp.start(vim.lsp.config.lua_ls)
				end,
			})

			vim.api.nvim_create_autocmd("FileType", {
				pattern = "go",
				callback = function()
					vim.lsp.start(vim.lsp.config.gopls)
				end,
			})

			-- Keymaps
			local opts = { noremap = true, silent = true }
			vim.keymap.set("n", "K", vim.lsp.buf.hover, opts)
			vim.keymap.set("n", "gd", vim.lsp.buf.definition, opts)
			vim.keymap.set({ "n", "v" }, "<leader>ca", vim.lsp.buf.code_action, opts)
		end,
	},
}
